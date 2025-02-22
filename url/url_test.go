/*
 * Copyright (c) 2021 Xingwang Liao
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package url

import (
	"testing"

	"github.com/hhq163/v8go"
)

func TestInject(t *testing.T) {
	t.Parallel()

	ctx := v8go.NewContext()

	if err := InjectTo(ctx); err != nil {
		t.Errorf("inject url polyfill: %v", err)
	}

	if val, _ := ctx.RunScript("typeof URL", ""); val.String() != "function" {
		t.Error("inject URL failed")
	}

	if val, _ := ctx.RunScript("typeof URLSearchParams", ""); val.String() != "function" {
		t.Error("inject URLSearchParams failed")
	}

	if val, _ := ctx.RunScript("new URLSearchParams('?a=1').get('a')", ""); val.String() != "1" {
		t.Error("test URLSearchParams failed")
	}
}
