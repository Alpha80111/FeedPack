#### template 

This is a template file which can be used to add new sources. Every new source should implement the following functions:

    FetchAndStoreFeedbacks(params models.Params, tenant string) ([]models.FeedbackIngest, error)
    IngestAndStoreFeedback(blob []byte, tenant string) (models.FeedbackIngest, error)

And the source.go GetProcessor function should be enhanced to return the particular feedback processor for the given key