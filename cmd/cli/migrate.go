package main

func doMigrate(arg2, arg3 string) error {
	dsn := getDSN()

	//run the migration command
	switch arg2 {
	case "up":
		err := bls.MigrateUp(dsn)
		if err != nil {
			return err
		}
	case "down":
		//migrate down --option 'all' get down all the previous migrations
		if arg3 == "all" {
			err := bls.MigrateDownAll(dsn)
			if err != nil {
				return err
			}
			//migrate down --option "" assuming get down the most recent migration only
		} else {
			err := bls.Steps(-1, dsn)
			if err != nil {
				return err
			}
		}
		//reset the entire database running down all migrations an then running up
	case "reset":
		//migrate down all the previous migrations
		err := bls.MigrateDownAll(dsn)
		if err != nil {
			return err
		}
		//migrate Up
		err = bls.MigrateUp(dsn)
		if err != nil {
			return err
		}
	//for the last case if something vet wrong
	default:
		showHelp()
	}
	return nil
}
