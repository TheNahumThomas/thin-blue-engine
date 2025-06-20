### Rule Engine Key Objectives

- Evaluates Data Inputs
- Matches Rules against Input Conditions and Actions
- Performs an action

Log -> Rule -> Evaluator -> Action Agenda -> Action Output

### Functional Requirements:

- The Engine shall accept ingress of both schema and schema-less JSON logs
- The Engine shall parse logs into custom data structures and types.

- The Engine shall provide pattern matching, rule evaluation, feature aggregation, regular expresions and optional expansions to provide high quality insights.

- The Engine shall be capable of "hot reloading" rule-sets without downtime.

- The engine shall perform pattern matching, rule evaluation, feature aggregation/ event thresholds, regular expressions and optional expansions for rules in a secure sandboxed envrionment.

- The Engine shall support rule linking and dynamic execution of those link outputs.

- The Engine shall dynamically prioritise actions sent to the alert worker in the event of potential flood/choke conditions.

- The engine shall feature inbuilt rule testing functionality

- The engine shall feature a straightforward, simple debug utility for the purpose of detecting false positives and unexpected rule behaviours.

### Non-Functional Requirements

- The engine should be optimised to run as a lightweight plug and play security component.

- The engine should produce consistently formatted, efficiently processed event objects as ingress for the dispatch worker.


log_structure -> interface
alert_object -> interface {
    UID
    Priority
    Type
    Set
    Insight
    TimeOffset
}

go provision_alert_queue{
    send index 0 to alert_dispatcher
    evaluate_priority() and reorder()
    
}

func load_rules

func match_rule

func match_aggregate

func match_regex

func match_pattern

func pattern_builder

func engine_async{
    go match_rule
    go match_aggregate
    go match_regex
    go match_pattern

   

}

func build_alert {
    event:
        quick_access_table_vals[eg, eg, eg, eg].query(QUERY)
    
    alert_object.new(MAP[QUERY])

    send_to(alert_dispatch_queue,priority)
     
}

