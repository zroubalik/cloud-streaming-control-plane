/*
Copyright 2020 The Knative Authors

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

package inbound

import (
	context "context"

	v1 "k8s.io/api/core/v1"
	reconciler "knative.dev/pkg/reconciler"

	v1alpha1 "knative.dev/streaming/pkg/apis/streaming/v1alpha1"
	inbound "knative.dev/streaming/pkg/client/injection/reconciler/streaming/v1alpha1/inbound"
)

// TODO: PLEASE COPY AND MODIFY THIS FILE AS A STARTING POINT

// newReconciledNormal makes a new reconciler event with event type Normal, and
// reason InboundReconciled.
func newReconciledNormal(namespace, name string) reconciler.Event {
	return reconciler.NewEvent(v1.EventTypeNormal, "InboundReconciled", "Inbound reconciled: \"%s/%s\"", namespace, name)
}

// Reconciler implements controller.Reconciler for Inbound resources.
type Reconciler struct {
	// TODO: add additional requirements here.
}

// Check that our Reconciler implements Interface
var _ inbound.Interface = (*Reconciler)(nil)

// Optionally check that our Reconciler implements Finalizer
//var _ inbound.Finalizer = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, o *v1alpha1.Inbound) reconciler.Event {
	o.Status.InitializeConditions()

	// TODO: add custom reconciliation logic here.

	o.Status.ObservedGeneration = o.Generation
	return newReconciledNormal(o.Namespace, o.Name)
}

// Optionally, use FinalizeKind to add finalizers. FinalizeKind will be called
// when the resource is deleted.
//func (r *Reconciler) FinalizeKind(ctx context.Context, o *v1alpha1.Inbound) reconciler.Event {
//	// TODO: add custom finalization logic here.
//	return nil
//}