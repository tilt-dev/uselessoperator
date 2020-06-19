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
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	tiltv1 "june18/api/v1"
)

// WebReconciler reconciles a Web object
type WebReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=tilt.op.tilt.dev,resources=webs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tilt.op.tilt.dev,resources=webs/status,verbs=get;update;patch

func (r *WebReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	log := r.Log.WithValues("web", req.NamespacedName)

	log.Info("hello from web ctrl", "name", req.NamespacedName)
	// your logic here

	return ctrl.Result{}, nil
}

func (r *WebReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tiltv1.Web{}).
		Watches(
			&source.Kind{Type: &tiltv1.Machine{}},
			&handler.EnqueueRequestsFromMapFunc{
				ToRequests: handler.ToRequestsFunc(
					func(obj handler.MapObject) []ctrl.Request {
						return []ctrl.Request{
							{
								NamespacedName: types.NamespacedName{
									Name:      obj.Meta.GetName(),
									Namespace: obj.Meta.GetNamespace(),
								},
							},
						}
					},
				)},
		).
		Complete(r)
}
