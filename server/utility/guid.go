package utility

/**
 * @author Dhawal Dyavanpalli <dhawalhost@gmail.com>
 * @desc Created on 2020-09-07 12:21:50 am
 * @copyright Crearosoft
 */

import (
	"github.com/segmentio/ksuid"
)

// GetGUID returns Unique GUID
func GetGUID() string {
	return ksuid.New().String()
}
