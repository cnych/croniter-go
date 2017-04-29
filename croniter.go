// croniter-go is a golang package to provide iteration for datetime object.

// Copyright 2017 YangChuanhu. icnych@gmail.com.
// All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// step_search_re = re.compile(r'^([^-]+)-([^-/]+)(/(.*))?$')
// search_re = re.compile(r'^([^-]+)-([^-/]+)(/(.*))?$')
// only_int_re = re.compile(r'^\d+$')
// any_int_re = re.compile(r'^\d+')
// star_or_int_re = re.compile(r'^(\d+|\*)$')
package croniter

import "time"

type CroniterBadCronError struct {
}

type CroniterBadDateError struct {
}

type CroniterNotAlphaError struct {
}

func (c *CroniterBadCronError) Error() string {
	return "croniter bad cron error"
}
func (c *CroniterBadDateError) Error() string {
	return "croniter bad date error"
}
func (c *CroniterNotAlphaError) Error() string {
	return "croniter not alpha error"
}

type Croniter struct {
	ExprFormat string
	StartTime  time.Time
}

func NewCroniter(expr string, startTime time.Time) (*Croniter, error) {
	if len(expr) != 5 || len(expr) != 6 {
		return nil, CroniterBadCronError{}
	}
	// TODO
	return &Croniter{
		ExprFormat: expr,
		StartTime:  startTime,
	}, nil
}

func (c *Croniter) Next() time.Time {
	return nil
}

func (c *Croniter) Prev() time.Time {
	return nil
}

func (c *Croniter) Current() time.Time {
	return nil
}
