package registry

import (
	"testing"

	kapi "github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/runtime"

	config "github.com/openshift/origin/pkg/config/api"
	template "github.com/openshift/origin/pkg/template/api"
)

func TestNewRESTInvalidType(t *testing.T) {
	storage := NewREST()
	_, err := storage.Create(nil, &kapi.Pod{})
	if err == nil {
		t.Errorf("Expected type error.")
	}
}

func TestNewRESTDefaultsName(t *testing.T) {
	storage := NewREST()
	obj, err := storage.Create(nil, &template.Template{
		ObjectMeta: kapi.ObjectMeta{
			Name: "test",
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, ok := obj.(*config.Config)
	if !ok {
		t.Fatalf("unexpected return object: %#v", obj)
	}
}

func TestNewRESTInvalidParameter(t *testing.T) {
	storage := NewREST()
	_, err := storage.Create(nil, &template.Template{
		ObjectMeta: kapi.ObjectMeta{
			Name: "test",
		},
		Parameters: []template.Parameter{
			{
				Name:     "TEST_PARAM",
				Generate: "[a-z0-Z0-9]{8}",
			},
		},
		Objects: []runtime.Object{},
	})
	if err == nil {
		t.Fatalf("Expected 'invalid parameter error', got nothing")
	}
}

func TestNewRESTTemplateLabels(t *testing.T) {
	testLabels := map[string]string{
		"label1": "value1",
		"label2": "value2",
	}
	storage := NewREST()
	obj, err := storage.Create(nil, &template.Template{
		ObjectMeta: kapi.ObjectMeta{
			Name: "test",
		},
		Objects: []runtime.Object{
			&kapi.Service{
				ObjectMeta: kapi.ObjectMeta{
					Name: "test-service",
				},
				Spec: kapi.ServiceSpec{
					Ports: []kapi.ServicePort{
						{
							Port:     80,
							Protocol: kapi.ProtocolTCP,
						},
					},
					SessionAffinity: kapi.AffinityTypeNone,
				},
			},
		},
		ObjectLabels: testLabels,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	config, ok := obj.(*config.Config)
	if !ok {
		t.Fatalf("unexpected return object: %#v", obj)
	}
	svc, ok := config.Items[0].(*kapi.Service)
	if !ok {
		t.Fatalf("Unexpected object in config: %#v", svc)
	}
	for k, v := range testLabels {
		value, ok := svc.Labels[k]
		if !ok {
			t.Fatalf("Missing output label: %s", k)
		}
		if value != v {
			t.Fatalf("Unexpected label value: %s", value)
		}
	}
}
