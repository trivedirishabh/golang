provider "atlan" {
  api_key = "your-api-key"
  base_url = "https://api.atlan.com"
}

resource "atlan_asset" "example" {
  name        = "Example Asset"
  description = "An example asset created using the custom Atlan provider."
}
