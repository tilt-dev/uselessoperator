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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	tiltv1 "op/api/v1"
)

// WebReconciler reconciles a Web object
type WebReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

var message string

// SetMessage sets the message
func SetMessage(s string) {
	message = s
}

// GetMessage returns the message
func GetMessage() string {
	return message
}

// +kubebuilder:rbac:groups=tilt.op.tilt.dev,resources=webs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tilt.op.tilt.dev,resources=webs/status,verbs=get;update;patch

func (r *WebReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

// https://twitter.com/ellenkorbes/status/1271248596055269376
func (r *WebReconciler) machineUpdate(obj handler.MapObject) []ctrl.Request {
	namespaced := types.NamespacedName{Namespace: obj.Meta.GetNamespace(), Name: obj.Meta.GetName()}
	ctx := context.Background()
	log := r.Log.WithValues("web", namespaced)
	var machine tiltv1.Machine
	if err := r.Get(ctx, namespaced, &machine); err != nil {
		log.Info("cant get machine", "name", namespaced)
		return []ctrl.Request{}
	}
	if machine.Spec.MachineType == "playful" {
		SetMessage(machine.Status.Status)
	}
	return []ctrl.Request{}
}

func (r *WebReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tiltv1.Web{}).
		Watches(
			&source.Kind{Type: &tiltv1.Machine{}},
			&handler.EnqueueRequestsFromMapFunc{
				ToRequests: handler.ToRequestsFunc(r.machineUpdate),
			}).
		Complete(r)
}
