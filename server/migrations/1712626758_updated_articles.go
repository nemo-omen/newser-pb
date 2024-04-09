package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("vvp5ozr9fg6j84w")
		if err != nil {
			return err
		}

		// update
		edit_site_link := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "u1u6h2an",
			"name": "site_link",
			"type": "url",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"exceptDomains": [],
				"onlyDomains": []
			}
		}`), edit_site_link); err != nil {
			return err
		}
		collection.Schema.AddField(edit_site_link)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("vvp5ozr9fg6j84w")
		if err != nil {
			return err
		}

		// update
		edit_site_link := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "u1u6h2an",
			"name": "feed_link",
			"type": "url",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"exceptDomains": [],
				"onlyDomains": []
			}
		}`), edit_site_link); err != nil {
			return err
		}
		collection.Schema.AddField(edit_site_link)

		return dao.SaveCollection(collection)
	})
}
