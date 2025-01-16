package behavior

import (
	"fmt"
	"strings"
)

type BehaviorManager struct {
	Personality string
	Tone        string
	CustomRules map[string]string
}

func NewBehaviorManager(personality, tone string) *BehaviorManager {
	return &BehaviorManager{
		Personality: personality,
		Tone:        tone,
		CustomRules: make(map[string]string),
	}
}

func (bm *BehaviorManager) SetPersonality(personality string) {
	bm.Personality = personality
	fmt.Printf("BehaviorManager: Personality set to '%s'\n", personality)
}

func (bm *BehaviorManager) SetTone(tone string) {
	bm.Tone = tone
	fmt.Printf("BehaviorManager: Tone set to '%s'\n", tone)
}

func (bm *BehaviorManager) AddCustomRule(trigger string, response string) {
	trigger = strings.ToLower(trigger)
	bm.CustomRules[trigger] = response
	fmt.Printf("BehaviorManager: Custom rule added for trigger '%s'\n", trigger)
}

func (bm *BehaviorManager) GetCustomResponse(input string) (string, bool) {
	input = strings.ToLower(input)
	response, exists := bm.CustomRules[input]
	if exists {
		fmt.Printf("BehaviorManager: Custom response triggered for input '%s'\n", input)
	}
	return response, exists
}

func (bm *BehaviorManager) AdjustResponse(input string) string {
	// Apply personality and tone adjustments to the response
	response := fmt.Sprintf("Personality: %s, Tone: %s, Input: %s", bm.Personality, bm.Tone, input)
	fmt.Printf("BehaviorManager: Adjusted response -> %s\n", response)
	return response
}