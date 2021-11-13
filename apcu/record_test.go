/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package apcu_test

import (
	"context"
	"testing"

	"github.com/go-spring/spring-base/apcu"
	"github.com/go-spring/spring-base/assert"
	"github.com/go-spring/spring-base/fastdev"
	"github.com/go-spring/spring-base/knife"
)

func TestRecord(t *testing.T) {

	fastdev.SetRecordMode(true)
	defer func() {
		fastdev.SetRecordMode(false)
	}()

	sessionID := fastdev.NewSessionID()
	ctx := knife.New(context.Background())
	err := knife.Set(ctx, fastdev.RecordSessionIDKey, sessionID)
	assert.Nil(t, err)

	defer func() {
		apcu.Delete(ctx, "a")
	}()

	type dataType struct {
		Data string `json:"a"`
	}

	var b *dataType
	ok, err := apcu.Load(ctx, "a", &b)
	assert.Nil(t, err)
	assert.False(t, ok)

	apcu.Store(ctx, "a", &dataType{
		Data: "success",
	})

	m := make(map[string]interface{})
	apcu.Range(func(key, value interface{}) bool {
		m[key.(string)] = value
		return true
	})
	assert.Equal(t, m["a"], &dataType{
		Data: "success",
	})

	ok, err = apcu.Load(ctx, "a", &b)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = apcu.Load(ctx, "a", &b)
	assert.Nil(t, err)
	assert.True(t, ok)

	fastdev.RecordInbound(ctx, &fastdev.Action{})
}
