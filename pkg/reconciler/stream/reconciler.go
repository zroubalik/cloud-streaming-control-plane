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

// Code generated by injection-gen. DO NOT EDIT.

package stream

import (
	context "context"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"knative.dev/pkg/logging"
	reconciler "knative.dev/pkg/reconciler"

	v1alpha1 "knative.dev/streaming/pkg/apis/streaming/v1alpha1"
	stream "knative.dev/streaming/pkg/client/injection/reconciler/streaming/v1alpha1/stream"
	"knative.dev/streaming/pkg/kafka"
)

var (
	HARDCODED_BOOTSTRAP_SERVERS = []string{"my-cluster-kafka-bootstrap.kafka.svc:9092"}
)

// newReconciledNormal makes a new reconciler event with event type Normal, and
// reason StreamReconciled.
func newReconciledNormal(namespace, name string) reconciler.Event {
	return reconciler.NewEvent(v1.EventTypeNormal, "StreamReconciled", "Stream reconciled: \"%s/%s\"", namespace, name)
}

// Reconciler implements controller.Reconciler for Stream resources.
type Reconciler struct {
	// Using a shared kafkaClusterAdmin does not work currently because of an issue with
	// Shopify/sarama, see https://github.com/Shopify/sarama/issues/1162.
	kafkaClusterAdmin sarama.ClusterAdmin
}

// Check that our Reconciler implements Interface
var _ stream.Interface = (*Reconciler)(nil)

// Optionally check that our Reconciler implements Finalizer
//var _ stream.Finalizer = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, stream *v1alpha1.Stream) reconciler.Event {
	o.Status.InitializeConditions()

	// TODO: add custom reconciliation logic here.

	kafkaClusterAdmin, err := r.createKafkaClusterClient(ctx)
	if err != nil {
		stream.Status.MarkConfigFailed("CannotConnectToKafka", "cannot create a connection to kafka cluster admin: %v", err)
		return err
	}

	stream.Status.MarkConfigTrue()

	if err := r.createTopic(ctx, stream, kafkaClusterAdmin); err != nil {
		stream.Status.MarkNotReady("TopicCreateFailed", "error while creating topic: %s", err)
		return err
	}

	stream.Status.MarkReady()

	o.Status.ObservedGeneration = o.Generation
	return newReconciledNormal(o.Namespace, o.Name)
}

func (r *Reconciler) createKafkaClusterClient(ctx context.Context) (sarama.ClusterAdmin, error) {
	// We don't currently initialize r.kafkaClusterAdmin, hence we end up creating the cluster admin client every time.
	// This is because of an issue with Shopify/sarama. See https://github.com/Shopify/sarama/issues/1162.
	// Once the issue is fixed we should use a shared cluster admin client. Also, r.kafkaClusterAdmin is currently
	// used to pass a fake admin client in the tests.
	kafkaClusterAdmin := r.kafkaClusterAdmin
	if kafkaClusterAdmin == nil {
		var err error
		kafkaClusterAdmin, err = kafka.MakeClusterAdminClient("knative-streaming-controller", HARDCODED_BOOTSTRAP_SERVERS)
		if err != nil {
			return nil, err
		}
	}
	return kafkaClusterAdmin, nil
}

func (r *Reconciler) createTopic(ctx context.Context, stream *v1alpha1.Stream, kafkaClusterAdmin sarama.ClusterAdmin) error {
	logger := logging.FromContext(ctx)

	topicName := kafka.TopicName(stream)
	logger.Info("Creating topic on Kafka cluster", zap.String("topic", topicName))
	err := kafkaClusterAdmin.CreateTopic(topicName, &sarama.TopicDetail{
		ReplicationFactor: 1,
		NumPartitions:     10,
	}, false)
	if e, ok := err.(*sarama.TopicError); ok && e.Err == sarama.ErrTopicAlreadyExists {
		return nil
	} else if err != nil {
		logger.Error("Error creating topic", zap.String("topic", topicName), zap.Error(err))
	} else {
		logger.Info("Successfully created topic", zap.String("topic", topicName))
	}
	return err
}
