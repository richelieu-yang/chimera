// Package nacosKit
/**
 * weight必须满足条件:	>= 0 && <= 10000
 * score必须满足条件: 	>= 0 && <= 100
 */
package nacosKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/floatKit"
)

// criticalScore
/**
 * score的临界点（默认30；Go和Java的值必须确保一样）.
 *
 * PS: 不能为public，
 */
var criticalScore float64 = 30

func checkScore(score float64) error {
	if score < 0 || score > 100 {
		return errorKit.New("score(%.2f) is invalid", score)
	}
	return nil
}

func checkWeight(weight float64) error {
	if weight < 0 || weight > 10000 {
		return errorKit.New("weight(%.2f) is invalid", weight)
	}
	return nil
}

func SetCriticalScore(score float64) error {
	if err := checkScore(score); err != nil {
		return err
	}

	criticalScore = score
	return nil
}

func GetCriticalScore() float64 {
	return criticalScore
}

// WeightToScore _
func WeightToScore(weight float64) (float64, error) {
	if err := checkWeight(weight); err != nil {
		return 0, err
	}

	if weight >= criticalScore*100 {
		return floatKit.Div(weight, 100), nil
	}
	return weight, nil
}

// ScoreToWeight _
func ScoreToWeight(score float64) (float64, error) {
	if err := checkScore(score); err != nil {
		return 0, err
	}

	if score >= criticalScore {
		return floatKit.Mul(score, 100), nil
	}
	return score, nil
}
