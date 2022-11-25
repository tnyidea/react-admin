# dataprovider

- getList
  - GET http://my.api.url/posts?_sort=title&_order=ASC&_start=0&_end=24
- getOne
  - GET http://my.api.url/posts/123
- getManyReference
  - GET http://my.api.url/posts?author_id=345
- getMany
  - GET http://my.api.url/posts?id=123&id=456&id=789
- create
  - POST http://my.api.url/posts/123
- update
  - PUT http://my.api.url/posts/123
- updateMany
  - PUT http://my.api.url/posts/123, PUT http://my.api.url/posts/456, PUT http://my.api.url/posts/789
- delete
  - DELETE http://my.api.url/posts/123