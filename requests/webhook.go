package requests

type Webhook struct {
	Action string `json:"action"`
	// Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// Just putting this here incase we find that we need more event info
//
// {
//   "action": "update",
//   "actor": {
//     "email": "user-0436@example.com",
//     "id": "096370c8-f3ad-4e20-a309-fb4f06ec0a89"
//   },
//   "created_at": "2016-10-26T22:50:14Z",
//   "id": "d472a8bb-1a3c-4f78-aad1-995e6d0022ec",
//   "data": {
//     "archived_at": null,
//     "buildpack_provided_description": "Ruby/Rails",
//     "build_stack": {
//       "id": "5e854079-da33-475b-a209-77f527c0e2fb",
//       "name": "cedar-14"
//     },
//     "created_at": "2016-10-26T22:50:14Z",
//     "id": "d4714cc8-aa56-4314-817c-0c6a66ff3d41",
//     "git_url": "https://git.heroku.com/sample-app-0301.git",
//     "maintenance": false,
//     "name": "sample-app-0301",
//     "owner": {
//       "email": "user-0436@example.com",
//       "id": "096370c8-f3ad-4e20-a309-fb4f06ec0a89"
//     },
//     "region": {
//       "id": "7044c05a-0873-42c2-abbe-6841c5481ba7",
//       "name": "us"
//     },
//     "organization": null,
//     "space": null,
//     "released_at": "2016-10-26T22:50:14Z",
//     "repo_size": 1048576,
//     "slug_size": null,
//     "stack": {
//       "id": "5e854079-da33-475b-a209-77f527c0e2fb",
//       "name": "cedar-14"
//     },
//     "updated_at": "2016-10-26T22:50:14Z",
//     "web_url": "https://sample-app-0301.herokuapp.com/"
//   },
//   "previous_data": {
//   },
//   "published_at": null,
//   "resource": "app",
//   "sequence": null,
//   "updated_at": "2016-10-26T22:50:14Z",
//   "version": "application/vnd.heroku+json; version=3",
//   "webhook_metadata": {
//     "attempt": {
//       "id": "8a44f820-2354-489d-9a11-a793cbf49979"
//     },
//     "delivery": {
//       "id": "d244009a-670f-4340-88e9-789a4f9002d5"
//     },
//     "event": {
//       "id": "d472a8bb-1a3c-4f78-aad1-995e6d0022ec",
//       "include": "api:app"
//     },
//     "webhook": {
//       "id": "01234567-89ab-cdef-0123-456789abcdef"
//     }
//   }
// }
