package handlers

import (
	"context"
	"testing"

	"github.com/machinebox/graphql"
)

func TestCreateKtp(t *testing.T) {
	// create a client (safe to share across requests)
	var client = graphql.NewClient("http://localhost:8080/query")

	var req = graphql.NewRequest(`
		mutation {
			createKtp(input : {
				nik:           "135",
				nama:          "ronda",
				jenis_kelamin: "male",
				tanggal_lahir: "2011-01-02 15:04:05",
				alamat:        "taman cita-cita no 2",
				agama:         "catholic"
			}){
				nama
				agama
			}
		}
	`)

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}
}

func TestGetAllKtp(t *testing.T) {
	// create a client (safe to share across requests)
	var client = graphql.NewClient("http://localhost:8080/query")

	var req = graphql.NewRequest(`
		query {
			ktp {
				nik
				nama
			}
		}
	`)

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}
}

func TestDeleteKtp(t *testing.T) {
	// create a client (safe to share across requests)
	var client = graphql.NewClient("http://localhost:8080/query")

	var req = graphql.NewRequest(`
		mutation {
			deleteKtp(id: 1)
		}
	`)

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}
}

func TestUpdateKtp(t *testing.T) {
	// create a client (safe to share across requests)
	var client = graphql.NewClient("http://localhost:8080/query")

	var req = graphql.NewRequest(`
		mutation {
			editKtp(id: 1, input : {
				nik:           "135",
				nama:          "ronda",
				jenis_kelamin: "male",
				tanggal_lahir: "2011-01-02 15:04:05",
				alamat:        "taman cita-cita no 2",
				agama:         "catholic",
			}){
				nama
				agama
			}
		}
	`)

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}
}
