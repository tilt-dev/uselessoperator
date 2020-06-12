/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	tiltv1 "op/api/v1"
)

// MachineReconciler reconciles a Machine object
type MachineReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=tilt.op.tilt.dev,resources=machines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tilt.op.tilt.dev,resources=machines/status,verbs=get;update;patch

func (r *MachineReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	rand.Seed(time.Now().UnixNano())
	sleep := func(t int) { time.Sleep(time.Duration(rand.Intn(t)) * time.Millisecond) }

	var machine tiltv1.Machine
	ctx := context.Background()
	log := r.Log.WithValues("machine", req.NamespacedName)
	if err := r.Get(ctx, req.NamespacedName, &machine); err != nil {
		if errors.IsNotFound(err) {
			log.Info("object not found", "name", req.NamespacedName)
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	if machine.Status.Status == "" {
		machine.Status.Status = "HOWDY"
		err := r.Status().Update(ctx, &machine)
		if err != nil {
			log.Error(err, "cant update status")
			return ctrl.Result{}, err
		}
		log.Info("howdy", "name", req.NamespacedName)
		return ctrl.Result{}, nil
	}
	if machine.Status.Status == "OK" {
		return ctrl.Result{}, nil
	}

	if machine.Status.Status == "DELETE" {
		sleep(5000)
		if err := r.Delete(ctx, &machine); err != nil {
			log.Error(err, "error deleting")
		}
		return ctrl.Result{}, nil
	}

	mtype := machine.Spec.MachineType
	switch mtype {
	case "useless":
		machine.Status.Status = "DELETE"
		log.Info("marked for deletion", "name", req.NamespacedName)
	case "useful":
		machine.Status.Status = "OK"
		temp := fmt.Sprint("machine allowed:", mtype)
		log.Info(temp)
	case "playful":

		var web tiltv1.Web
		if err := r.Get(ctx, types.NamespacedName{Namespace: req.Namespace, Name: req.Name + "-web"}, &web); err != nil {
			if errors.IsNotFound(err) {
				log.Info("object not found, creating", "name", req.NamespacedName)
				if err := r.Create(ctx, r.newWeb(machine)); err != nil {
					log.Error(err, "cant create web")
					return ctrl.Result{}, err
				}
				return ctrl.Result{}, nil
			} else {
				log.Error(err, "get issue")
				return ctrl.Result{}, err
			}
		}

		r := "o"
		s := `\`
		if len(machine.Status.Status) > 10 {
			machine.Status.Status = strings.Trim(machine.Status.Status[0:10], " ")
		}
		if machine.Status.Status == "HOWDY" {
			machine.Status.Status = r
			break
		}
		if machine.Status.Status == fill(10, r) {
			time.Sleep(time.Second)
			machine.Status.Status = "DELETE"
			break
		}
		sleep(500)
		machine.Status.Status = fill(plusminus(len(machine.Status.Status)), r)
		machine.Status.Status = machine.Status.Status + fill(10-len(machine.Status.Status), " ") + s
	}
	if err := r.Status().Update(ctx, &machine); err != nil {
		log.Error(err, "cant update status")
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func fill(n int, c string) string {
	r := ""
	for i := 0; i < n; i++ {
		r += c
	}
	return r
}

func plusminus(count int) int {
	if count == 0 {
		return 1
	}
	if count == 1 {
		return 2
	}
	n := rand.Intn(2)
	if n == 0 {
		n = -1
	}
	n += count
	return n
}

func (r *MachineReconciler) newWeb(machine tiltv1.Machine) *tiltv1.Web {
	return &tiltv1.Web{
		TypeMeta: metav1.TypeMeta{Kind: "Web"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      machine.Name + "-web",
			Namespace: machine.Namespace,
		},
		Spec: tiltv1.WebSpec{},
	}
}

func (r *MachineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tiltv1.Machine{}).
		Owns(&tiltv1.Web{}).
		Complete(r)
}
