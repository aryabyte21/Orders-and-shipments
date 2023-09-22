```
curl -X PATCH -H "Content-Type: application/json" -d '{
  "status": "shipped"
}' "<http://localhost:8080/updateShipmentStatus?shipmentID=650de5147280cff5781678bd>"
```