package cli

import (
	"context"
	"errors"
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kfake "k8s.io/client-go/kubernetes/fake"
)

func Test_waitForPodStatus(t *testing.T) {

	t.Run("successfully errors on cancelled context", func(t *testing.T) {
		k := &KubeClient{
			cli:        kfake.NewSimpleClientset(),
			instanceID: "test",
		}

		podSpec := &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: defaultNamespace},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{Name: "test-pod", Image: "containous/whoami"},
				},
			},
		}

		ctx, cancel := context.WithCancel(context.TODO())
		cancel()
		err := k.waitForPodStatus(ctx, v1.PodRunning, podSpec)
		if !errors.Is(err, context.Canceled) {
			t.Errorf("waitForPodStatus should throw context cancellation error; err=%s", err)
		}
	})

	t.Run("successfully errors on timeout", func(t *testing.T) {
		k := &KubeClient{
			cli:        kfake.NewSimpleClientset(),
			instanceID: "test",
		}

		podSpec := &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: defaultNamespace},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{Name: "test-pod", Image: "containous/whoami"},
				},
			},
		}

		pod, err := k.cli.CoreV1().Pods(defaultNamespace).Create(context.Background(), podSpec, metav1.CreateOptions{})
		if err != nil {
			t.Errorf("failed to create pod; err=%s", err)
		}
		defer k.cli.CoreV1().Pods(defaultNamespace).Delete(context.Background(), pod.Name, metav1.DeleteOptions{})

		ctx, cancelFunc := context.WithTimeout(context.TODO(), 0*time.Second)
		defer cancelFunc()
		err = k.waitForPodStatus(ctx, v1.PodRunning, podSpec)
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("waitForPodStatus should throw deadline exceeded error; err=%s", err)
		}
	})

}
