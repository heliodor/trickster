/*
 * Copyright 2018 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package registration

import (
	"testing"

	"github.com/tricksterproxy/trickster/pkg/config"
	"github.com/tricksterproxy/trickster/pkg/tracing/options"
	tl "github.com/tricksterproxy/trickster/pkg/util/log"
)

func TestRegisterAll(t *testing.T) {

	// test nil config
	f, err := RegisterAll(nil, tl.ConsoleLogger("error"), true)
	if err == nil {
		t.Error("expected error for no config provided")
	}
	if len(f) > 0 {
		t.Errorf("expected %d got %d", 0, len(f))
	}

	// test good config
	f, err = RegisterAll(config.NewConfig(), tl.ConsoleLogger("error"), true)
	if err != nil {
		t.Error(err)
	}
	if len(f) != 1 {
		t.Errorf("expected %d got %d", 1, len(f))
	}

	// test bad implementation
	cfg := config.NewConfig()
	tc := options.NewOptions()

	cfg.TracingConfigs = make(map[string]*options.Options)
	cfg.TracingConfigs["test"] = tc
	cfg.TracingConfigs["test3"] = tc
	cfg.Origins["default"].TracingConfigName = "test"

	_, err = RegisterAll(cfg, tl.ConsoleLogger("error"), true)
	if err != nil {
		t.Error(err)
	}

	tc.TracerType = "jaeger"
	tc.CollectorURL = "http://example.com"
	_, err = RegisterAll(cfg, tl.ConsoleLogger("error"), false)
	if err != nil {
		t.Error(err)
	}

	tc.TracerType = "stdout"
	_, err = RegisterAll(cfg, tl.ConsoleLogger("error"), true)
	if err != nil {
		t.Error(err)
	}

	tc.TracerType = "zipkin"
	_, err = RegisterAll(cfg, tl.ConsoleLogger("error"), true)
	if err != nil {
		t.Error(err)
	}

	tc.TracerType = "foo"

	_, err = RegisterAll(cfg, tl.ConsoleLogger("error"), true)
	if err == nil {
		t.Error("expected error for invalid tracer type")
	}

	// test empty implementation
	tc.TracerType = ""
	f, _ = RegisterAll(cfg, tl.ConsoleLogger("error"), true)
	if len(f) > 0 {
		t.Errorf("expected %d got %d", 0, len(f))
	}

	tc.TracerType = "none"
	cfg.Origins["default"].TracingConfigName = "test2"
	_, err = RegisterAll(cfg, tl.ConsoleLogger("error"), true)
	if err == nil {
		t.Error("expected error for invalid tracing config name")
	}
	cfg.Origins["default"].TracingConfigName = "test"

	temp := cfg.TracingConfigs
	cfg.TracingConfigs = nil
	// test nil tracing config
	f, _ = RegisterAll(cfg, tl.ConsoleLogger("error"), true)
	if len(f) > 0 {
		t.Errorf("expected %d got %d", 0, len(f))
	}
	cfg.TracingConfigs = temp

	// test nil origin config
	cfg.Origins = nil
	_, err = RegisterAll(cfg, tl.ConsoleLogger("error"), true)
	if err == nil {
		t.Error("expected error for invalid tracing implementation")
	}

}

func TestGetTracer(t *testing.T) {
	tr, _ := GetTracer(nil, tl.ConsoleLogger("error"), true)
	if tr != nil {
		t.Error("expected nil tracer")
	}
}
