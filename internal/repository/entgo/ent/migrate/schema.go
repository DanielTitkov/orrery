// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// InterpretationsColumns holds the columns for the "interpretations" table.
	InterpretationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "range", Type: field.TypeJSON},
		{Name: "scale_interpretations", Type: field.TypeUUID, Nullable: true},
	}
	// InterpretationsTable holds the schema information for the "interpretations" table.
	InterpretationsTable = &schema.Table{
		Name:       "interpretations",
		Columns:    InterpretationsColumns,
		PrimaryKey: []*schema.Column{InterpretationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "interpretations_scales_interpretations",
				Columns:    []*schema.Column{InterpretationsColumns[4]},
				RefColumns: []*schema.Column{ScalesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// InterpretationTranslationsColumns holds the columns for the "interpretation_translations" table.
	InterpretationTranslationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "locale", Type: field.TypeEnum, Enums: []string{"en", "ru"}},
		{Name: "content", Type: field.TypeString},
		{Name: "interpretation_translations", Type: field.TypeUUID, Nullable: true},
	}
	// InterpretationTranslationsTable holds the schema information for the "interpretation_translations" table.
	InterpretationTranslationsTable = &schema.Table{
		Name:       "interpretation_translations",
		Columns:    InterpretationTranslationsColumns,
		PrimaryKey: []*schema.Column{InterpretationTranslationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "interpretation_translations_interpretations_translations",
				Columns:    []*schema.Column{InterpretationTranslationsColumns[3]},
				RefColumns: []*schema.Column{InterpretationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "interpretationtranslation_locale_interpretation_translations",
				Unique:  true,
				Columns: []*schema.Column{InterpretationTranslationsColumns[1], InterpretationTranslationsColumns[3]},
			},
		},
	}
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "steps", Type: field.TypeInt, Default: 2},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:       "items",
		Columns:    ItemsColumns,
		PrimaryKey: []*schema.Column{ItemsColumns[0]},
	}
	// ItemTranslationsColumns holds the columns for the "item_translations" table.
	ItemTranslationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "locale", Type: field.TypeEnum, Enums: []string{"en", "ru"}},
		{Name: "content", Type: field.TypeString},
		{Name: "item_translations", Type: field.TypeUUID, Nullable: true},
	}
	// ItemTranslationsTable holds the schema information for the "item_translations" table.
	ItemTranslationsTable = &schema.Table{
		Name:       "item_translations",
		Columns:    ItemTranslationsColumns,
		PrimaryKey: []*schema.Column{ItemTranslationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "item_translations_items_translations",
				Columns:    []*schema.Column{ItemTranslationsColumns[3]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "itemtranslation_locale_item_translations",
				Unique:  true,
				Columns: []*schema.Column{ItemTranslationsColumns[1], ItemTranslationsColumns[3]},
			},
		},
	}
	// NormsColumns holds the columns for the "norms" table.
	NormsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "base", Type: field.TypeInt, Default: 0},
		{Name: "mean", Type: field.TypeFloat64},
		{Name: "sigma", Type: field.TypeFloat64},
		{Name: "rank", Type: field.TypeInt, Default: 0},
		{Name: "meta", Type: field.TypeJSON, Nullable: true},
		{Name: "sample_norms", Type: field.TypeUUID},
		{Name: "scale_norms", Type: field.TypeUUID},
	}
	// NormsTable holds the schema information for the "norms" table.
	NormsTable = &schema.Table{
		Name:       "norms",
		Columns:    NormsColumns,
		PrimaryKey: []*schema.Column{NormsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "norms_samples_norms",
				Columns:    []*schema.Column{NormsColumns[9]},
				RefColumns: []*schema.Column{SamplesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "norms_scales_norms",
				Columns:    []*schema.Column{NormsColumns[10]},
				RefColumns: []*schema.Column{ScalesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "norm_sample_norms_scale_norms",
				Unique:  true,
				Columns: []*schema.Column{NormsColumns[9], NormsColumns[10]},
			},
		},
	}
	// QuestionsColumns holds the columns for the "questions" table.
	QuestionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "order", Type: field.TypeInt, Default: 10},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"simple"}, Default: "simple"},
	}
	// QuestionsTable holds the schema information for the "questions" table.
	QuestionsTable = &schema.Table{
		Name:       "questions",
		Columns:    QuestionsColumns,
		PrimaryKey: []*schema.Column{QuestionsColumns[0]},
	}
	// QuestionTranslationsColumns holds the columns for the "question_translations" table.
	QuestionTranslationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "locale", Type: field.TypeEnum, Enums: []string{"en", "ru"}},
		{Name: "content", Type: field.TypeString, Nullable: true},
		{Name: "header_content", Type: field.TypeString, Nullable: true},
		{Name: "footer_content", Type: field.TypeString, Nullable: true},
		{Name: "question_translations", Type: field.TypeUUID, Nullable: true},
	}
	// QuestionTranslationsTable holds the schema information for the "question_translations" table.
	QuestionTranslationsTable = &schema.Table{
		Name:       "question_translations",
		Columns:    QuestionTranslationsColumns,
		PrimaryKey: []*schema.Column{QuestionTranslationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "question_translations_questions_translations",
				Columns:    []*schema.Column{QuestionTranslationsColumns[5]},
				RefColumns: []*schema.Column{QuestionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "questiontranslation_locale_question_translations",
				Unique:  true,
				Columns: []*schema.Column{QuestionTranslationsColumns[1], QuestionTranslationsColumns[5]},
			},
		},
	}
	// ResponsesColumns holds the columns for the "responses" table.
	ResponsesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "value", Type: field.TypeInt, Default: 0},
		{Name: "meta", Type: field.TypeJSON, Nullable: true},
		{Name: "item_responses", Type: field.TypeUUID},
		{Name: "take_responses", Type: field.TypeUUID},
	}
	// ResponsesTable holds the schema information for the "responses" table.
	ResponsesTable = &schema.Table{
		Name:       "responses",
		Columns:    ResponsesColumns,
		PrimaryKey: []*schema.Column{ResponsesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "responses_items_responses",
				Columns:    []*schema.Column{ResponsesColumns[5]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "responses_takes_responses",
				Columns:    []*schema.Column{ResponsesColumns[6]},
				RefColumns: []*schema.Column{TakesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "response_item_responses_take_responses",
				Unique:  true,
				Columns: []*schema.Column{ResponsesColumns[5], ResponsesColumns[6]},
			},
		},
	}
	// ResultsColumns holds the columns for the "results" table.
	ResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "raw_score", Type: field.TypeFloat64},
		{Name: "final_score", Type: field.TypeFloat64},
		{Name: "meta", Type: field.TypeJSON, Nullable: true},
		{Name: "scale_results", Type: field.TypeUUID},
		{Name: "take_results", Type: field.TypeUUID},
	}
	// ResultsTable holds the schema information for the "results" table.
	ResultsTable = &schema.Table{
		Name:       "results",
		Columns:    ResultsColumns,
		PrimaryKey: []*schema.Column{ResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "results_scales_results",
				Columns:    []*schema.Column{ResultsColumns[6]},
				RefColumns: []*schema.Column{ScalesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "results_takes_results",
				Columns:    []*schema.Column{ResultsColumns[7]},
				RefColumns: []*schema.Column{TakesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "result_scale_results_take_results",
				Unique:  true,
				Columns: []*schema.Column{ResultsColumns[6], ResultsColumns[7]},
			},
		},
	}
	// SamplesColumns holds the columns for the "samples" table.
	SamplesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "criteria", Type: field.TypeJSON},
	}
	// SamplesTable holds the schema information for the "samples" table.
	SamplesTable = &schema.Table{
		Name:       "samples",
		Columns:    SamplesColumns,
		PrimaryKey: []*schema.Column{SamplesColumns[0]},
	}
	// ScalesColumns holds the columns for the "scales" table.
	ScalesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "global", Type: field.TypeBool, Default: false},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"sten", "mean", "sum", "zscore"}, Default: "sten"},
	}
	// ScalesTable holds the schema information for the "scales" table.
	ScalesTable = &schema.Table{
		Name:       "scales",
		Columns:    ScalesColumns,
		PrimaryKey: []*schema.Column{ScalesColumns[0]},
	}
	// ScaleItemsColumns holds the columns for the "scale_items" table.
	ScaleItemsColumns = []*schema.Column{
		{Name: "reverse", Type: field.TypeBool, Default: false},
		{Name: "item_id", Type: field.TypeUUID},
		{Name: "scale_id", Type: field.TypeUUID},
	}
	// ScaleItemsTable holds the schema information for the "scale_items" table.
	ScaleItemsTable = &schema.Table{
		Name:       "scale_items",
		Columns:    ScaleItemsColumns,
		PrimaryKey: []*schema.Column{},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "scale_items_items_item",
				Columns:    []*schema.Column{ScaleItemsColumns[1]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "scale_items_scales_scale",
				Columns:    []*schema.Column{ScaleItemsColumns[2]},
				RefColumns: []*schema.Column{ScalesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ScaleTranslationsColumns holds the columns for the "scale_translations" table.
	ScaleTranslationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "locale", Type: field.TypeEnum, Enums: []string{"en", "ru"}},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "scale_translations", Type: field.TypeUUID, Nullable: true},
	}
	// ScaleTranslationsTable holds the schema information for the "scale_translations" table.
	ScaleTranslationsTable = &schema.Table{
		Name:       "scale_translations",
		Columns:    ScaleTranslationsColumns,
		PrimaryKey: []*schema.Column{ScaleTranslationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "scale_translations_scales_translations",
				Columns:    []*schema.Column{ScaleTranslationsColumns[4]},
				RefColumns: []*schema.Column{ScalesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "scaletranslation_locale_scale_translations",
				Unique:  true,
				Columns: []*schema.Column{ScaleTranslationsColumns[1], ScaleTranslationsColumns[4]},
			},
		},
	}
	// TakesColumns holds the columns for the "takes" table.
	TakesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "seed", Type: field.TypeInt64, Default: 1656021440},
		{Name: "progress", Type: field.TypeInt, Default: 0},
		{Name: "page", Type: field.TypeInt, Default: 0},
		{Name: "start_time", Type: field.TypeTime, Nullable: true},
		{Name: "end_time", Type: field.TypeTime, Nullable: true},
		{Name: "suspicious", Type: field.TypeBool, Default: false},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"intro", "questions", "finish", "result"}, Default: "intro"},
		{Name: "meta", Type: field.TypeJSON, Nullable: true},
		{Name: "test_takes", Type: field.TypeUUID},
		{Name: "user_takes", Type: field.TypeUUID},
	}
	// TakesTable holds the schema information for the "takes" table.
	TakesTable = &schema.Table{
		Name:       "takes",
		Columns:    TakesColumns,
		PrimaryKey: []*schema.Column{TakesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "takes_tests_takes",
				Columns:    []*schema.Column{TakesColumns[11]},
				RefColumns: []*schema.Column{TestsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "takes_users_takes",
				Columns:    []*schema.Column{TakesColumns[12]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TestsColumns holds the columns for the "tests" table.
	TestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString, Unique: true, Size: 100},
		{Name: "published", Type: field.TypeBool, Default: true},
		{Name: "available_locales", Type: field.TypeJSON, Nullable: true},
	}
	// TestsTable holds the schema information for the "tests" table.
	TestsTable = &schema.Table{
		Name:       "tests",
		Columns:    TestsColumns,
		PrimaryKey: []*schema.Column{TestsColumns[0]},
	}
	// TestDisplaysColumns holds the columns for the "test_displays" table.
	TestDisplaysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "randomize_order", Type: field.TypeBool, Default: false},
		{Name: "questions_per_page", Type: field.TypeInt, Default: 1},
		{Name: "test_display", Type: field.TypeUUID, Unique: true},
	}
	// TestDisplaysTable holds the schema information for the "test_displays" table.
	TestDisplaysTable = &schema.Table{
		Name:       "test_displays",
		Columns:    TestDisplaysColumns,
		PrimaryKey: []*schema.Column{TestDisplaysColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "test_displays_tests_display",
				Columns:    []*schema.Column{TestDisplaysColumns[3]},
				RefColumns: []*schema.Column{TestsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TestTranslationsColumns holds the columns for the "test_translations" table.
	TestTranslationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "locale", Type: field.TypeEnum, Enums: []string{"en", "ru"}},
		{Name: "title", Type: field.TypeString, Size: 140},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "instruction", Type: field.TypeString, Nullable: true},
		{Name: "test_translations", Type: field.TypeUUID, Nullable: true},
	}
	// TestTranslationsTable holds the schema information for the "test_translations" table.
	TestTranslationsTable = &schema.Table{
		Name:       "test_translations",
		Columns:    TestTranslationsColumns,
		PrimaryKey: []*schema.Column{TestTranslationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "test_translations_tests_translations",
				Columns:    []*schema.Column{TestTranslationsColumns[5]},
				RefColumns: []*schema.Column{TestsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "testtranslation_locale_test_translations",
				Unique:  true,
				Columns: []*schema.Column{TestTranslationsColumns[1], TestTranslationsColumns[5]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "locale", Type: field.TypeEnum, Enums: []string{"en", "ru"}},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "picture", Type: field.TypeString, Nullable: true, Default: "https://www.gravatar.com/avatar/00000000000000000000000000000000?d=mp&f=y"},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "admin", Type: field.TypeBool, Default: false},
		{Name: "anonymous", Type: field.TypeBool, Default: false},
		{Name: "meta", Type: field.TypeJSON, Nullable: true},
		{Name: "user_aliases", Type: field.TypeUUID, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_users_aliases",
				Columns:    []*schema.Column{UsersColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserSessionsColumns holds the columns for the "user_sessions" table.
	UserSessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "sid", Type: field.TypeString, Unique: true},
		{Name: "ip", Type: field.TypeString},
		{Name: "user_agent", Type: field.TypeString},
		{Name: "last_activity", Type: field.TypeTime},
		{Name: "active", Type: field.TypeBool, Default: false},
		{Name: "meta", Type: field.TypeJSON, Nullable: true},
		{Name: "user_sessions", Type: field.TypeUUID},
	}
	// UserSessionsTable holds the schema information for the "user_sessions" table.
	UserSessionsTable = &schema.Table{
		Name:       "user_sessions",
		Columns:    UserSessionsColumns,
		PrimaryKey: []*schema.Column{UserSessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_sessions_users_sessions",
				Columns:    []*schema.Column{UserSessionsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "usersession_active",
				Unique:  false,
				Columns: []*schema.Column{UserSessionsColumns[7]},
			},
		},
	}
	// QuestionItemsColumns holds the columns for the "question_items" table.
	QuestionItemsColumns = []*schema.Column{
		{Name: "question_id", Type: field.TypeUUID},
		{Name: "item_id", Type: field.TypeUUID},
	}
	// QuestionItemsTable holds the schema information for the "question_items" table.
	QuestionItemsTable = &schema.Table{
		Name:       "question_items",
		Columns:    QuestionItemsColumns,
		PrimaryKey: []*schema.Column{QuestionItemsColumns[0], QuestionItemsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "question_items_question_id",
				Columns:    []*schema.Column{QuestionItemsColumns[0]},
				RefColumns: []*schema.Column{QuestionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "question_items_item_id",
				Columns:    []*schema.Column{QuestionItemsColumns[1]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TestQuestionsColumns holds the columns for the "test_questions" table.
	TestQuestionsColumns = []*schema.Column{
		{Name: "test_id", Type: field.TypeUUID},
		{Name: "question_id", Type: field.TypeUUID},
	}
	// TestQuestionsTable holds the schema information for the "test_questions" table.
	TestQuestionsTable = &schema.Table{
		Name:       "test_questions",
		Columns:    TestQuestionsColumns,
		PrimaryKey: []*schema.Column{TestQuestionsColumns[0], TestQuestionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "test_questions_test_id",
				Columns:    []*schema.Column{TestQuestionsColumns[0]},
				RefColumns: []*schema.Column{TestsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "test_questions_question_id",
				Columns:    []*schema.Column{TestQuestionsColumns[1]},
				RefColumns: []*schema.Column{QuestionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TestScalesColumns holds the columns for the "test_scales" table.
	TestScalesColumns = []*schema.Column{
		{Name: "test_id", Type: field.TypeUUID},
		{Name: "scale_id", Type: field.TypeUUID},
	}
	// TestScalesTable holds the schema information for the "test_scales" table.
	TestScalesTable = &schema.Table{
		Name:       "test_scales",
		Columns:    TestScalesColumns,
		PrimaryKey: []*schema.Column{TestScalesColumns[0], TestScalesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "test_scales_test_id",
				Columns:    []*schema.Column{TestScalesColumns[0]},
				RefColumns: []*schema.Column{TestsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "test_scales_scale_id",
				Columns:    []*schema.Column{TestScalesColumns[1]},
				RefColumns: []*schema.Column{ScalesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		InterpretationsTable,
		InterpretationTranslationsTable,
		ItemsTable,
		ItemTranslationsTable,
		NormsTable,
		QuestionsTable,
		QuestionTranslationsTable,
		ResponsesTable,
		ResultsTable,
		SamplesTable,
		ScalesTable,
		ScaleItemsTable,
		ScaleTranslationsTable,
		TakesTable,
		TestsTable,
		TestDisplaysTable,
		TestTranslationsTable,
		UsersTable,
		UserSessionsTable,
		QuestionItemsTable,
		TestQuestionsTable,
		TestScalesTable,
	}
)

func init() {
	InterpretationsTable.ForeignKeys[0].RefTable = ScalesTable
	InterpretationTranslationsTable.ForeignKeys[0].RefTable = InterpretationsTable
	ItemTranslationsTable.ForeignKeys[0].RefTable = ItemsTable
	NormsTable.ForeignKeys[0].RefTable = SamplesTable
	NormsTable.ForeignKeys[1].RefTable = ScalesTable
	QuestionTranslationsTable.ForeignKeys[0].RefTable = QuestionsTable
	ResponsesTable.ForeignKeys[0].RefTable = ItemsTable
	ResponsesTable.ForeignKeys[1].RefTable = TakesTable
	ResultsTable.ForeignKeys[0].RefTable = ScalesTable
	ResultsTable.ForeignKeys[1].RefTable = TakesTable
	ScaleItemsTable.ForeignKeys[0].RefTable = ItemsTable
	ScaleItemsTable.ForeignKeys[1].RefTable = ScalesTable
	ScaleTranslationsTable.ForeignKeys[0].RefTable = ScalesTable
	TakesTable.ForeignKeys[0].RefTable = TestsTable
	TakesTable.ForeignKeys[1].RefTable = UsersTable
	TestDisplaysTable.ForeignKeys[0].RefTable = TestsTable
	TestTranslationsTable.ForeignKeys[0].RefTable = TestsTable
	UsersTable.ForeignKeys[0].RefTable = UsersTable
	UserSessionsTable.ForeignKeys[0].RefTable = UsersTable
	QuestionItemsTable.ForeignKeys[0].RefTable = QuestionsTable
	QuestionItemsTable.ForeignKeys[1].RefTable = ItemsTable
	TestQuestionsTable.ForeignKeys[0].RefTable = TestsTable
	TestQuestionsTable.ForeignKeys[1].RefTable = QuestionsTable
	TestScalesTable.ForeignKeys[0].RefTable = TestsTable
	TestScalesTable.ForeignKeys[1].RefTable = ScalesTable
}
