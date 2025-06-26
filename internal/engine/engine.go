package engine

// Pseudo-code for Go-based Rule Engine Functional Requirements

// 1. Accept ingress of both schema and schema-less JSON logs


// 3. Pattern matching, rule evaluation, feature aggregation, regex, expansions
func ProcessLog(logObj LogObject) {
    features := AggregateFeatures(logObj)
    for rule in RuleSet {
        if MatchPattern(rule.Pattern, features) &&
           EvaluateRule(rule, features) &&
           MatchRegex(rule.Regex, features) {
            insights := ExpandRule(rule, features)
            SendToAgenda(insights)
        }
    }
}

// 4. Hot reloading rule-sets without downtime
func WatchRuleSet() {
    for {
        if RuleSetChanged() {
            ReloadRuleSet()
        }
        Sleep(Interval)
    }
}

// 5. Secure sandboxed rule evaluation
func EvaluateRule(rule Rule, features Features) bool {
    // Run rule logic in sandboxed environment (e.g., plugin, VM, or restricted goroutine)
    return wasmEvaluate(event_pointer, event_size)
}

// 6. Rule linking and dynamic execution of link outputs
func ExpandRule(rule Rule, features Features) []Insight {
    insights := []
    for link in rule.LinkedRules {
        output := ExecuteRule(link, features)
        insights = append(insights, output)
    }
    return insights
}

// 7. Dynamically prioritise actions sent to alert worker
func SendToAgenda(insights []Insight) {
    for insight in insights {
        priority := CalculatePriority(insight)
        AlertQueue.Enqueue(insight, priority)
    }
    AlertQueue.ReorderIfFlood()
}

// 8. Inbuilt rule testing functionality
func TestRule(rule Rule, testLog LogObject) TestResult {
    result := EvaluateRule(rule, AggregateFeatures(testLog))
    return CompareWithExpected(result, rule.ExpectedOutcome)
}

// 9. Simple debug utility for false positives/unexpected behaviours
func DebugRule(rule Rule, logObj LogObject) DebugInfo {
    trace := TraceRuleEvaluation(rule, logObj)
    if trace.TriggeredUnexpectedly {
        LogDebug(trace)
    }
    return trace
}