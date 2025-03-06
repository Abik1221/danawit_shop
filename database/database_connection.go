package database 

func DbInstance () mongo.Client{
	db, err := mongo.NewClient(options.Client().applyURI("MONGODB_URL"))
    if err != nil{
         log.Fatal(err)
	}

	var ctx, cancel = context.WithTimeout(context.Bckground(), 10 * time.Second)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
