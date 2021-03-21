# Guide
- https://echo.labstack.com/guide

# hello-echo-world
Web API by Echo and SQLite

# Rails Schema

```ruby
  create_table "fake_materials", force: :cascade do |t|
    t.string "name"
    t.string "email"
    t.string "credit_card_number"
    t.date "credit_card_expiry_date"
    t.string "credit_card_type"
    t.string "city"
    t.string "gender"
    t.integer "my_number"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
  end
```
