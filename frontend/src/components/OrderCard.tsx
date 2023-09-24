import React, { useState } from "react";
import {
  Card,
  Text,
  Switch,
  Collapse,
  Button,
  Badge,
  Divider,
  Space,
  Paper,
} from "@mantine/core";
import {
  IconHome2,
  IconDeviceMobile,
  IconClockHour5,
  IconTags,
  IconBrandStripe,
} from "@tabler/icons-react";
import { notifications } from "@mantine/notifications";

type Product = {
  Name: string;
  Price: string;
};

type Order = {
  OrderTotal: string;
  OrderNumber: number;
  CustomerName: Customer;
  CustomerEmail: string;
  CustomerOrderNumber: string;
  CustomerPhone: string;
  CustomerAddress: string;
  ProviderOrderId: number;
  CreatedAt: string;
  ProviderCreatedAt: string;
  Products: Product[];
  UniqueId: string;
  ID: number;
  stripe_payment_id: string;
};

type Customer = {
  FirstName: string;
  LastName: string;
  Currency: string;
  DefaultAddress: Address;
};

type Address = {
  Phone: string;
  Address1: string;
  City: string;
  Province: string;
  Zip: string;
  Country: string;
};
interface OrderCardProps {
  order: Order;
}

const OrderCard: React.FC<OrderCardProps> = ({ order }) => {
  const [expanded, setExpanded] = useState(false);
  const [status, setStatus] = useState("pending");
  console.log(order);
  const handleStatusChange = (checked: boolean) => {
    setStatus(checked ? "shipped" : "pending");
    fetch(`http://localhost:8080/updateShipmentStatus?shipmentID=${order.ID}`, {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        status: checked ? "shipped" : "pending",
      }),
    }).then((res) => {
      if (res.status === 200) {
        notifications.show({
          title: "Shipment status updated",
          message: `Shipment status updated to ${
            checked ? "shipped" : "pending"
          }`,
          color: checked ? "green" : "yellow",
          autoClose: 3000,
        });
      }
    });
  };

  return (
    <Card
      shadow="sm"
      padding="lg"
      radius="md"
      withBorder
      style={{ margin: "1rem" }}
    >
      <Text
        size="xl"
        fw={900}
        variant="gradient"
        gradient={{ from: "blue", to: "teal", deg: 186 }}
      >
        {order.CustomerName.FirstName} {order.CustomerName.LastName}
      </Text>
      <Text c="dimmed">{order.CustomerEmail}</Text>

      <div
        style={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          fontSize: "0.9rem",
          padding: "0.5rem",
        }}
      >
        <Badge variant="light" leftSection="💵">
          {order.OrderTotal} {order.CustomerName.Currency}
        </Badge>
      </div>
      <div style={{ display: "flex", alignItems: "flex-start" }}>
        <Switch
          checked={status === "shipped"}
          onChange={(e) => handleStatusChange(e.target.checked)}
          style={{ marginRight: "0.5rem" }}
        />
        <Text style={{ paddingBottom: "0.7rem" }} size="sm">
          Shipment Status
        </Text>
      </div>
      <Button
        variant="light"
        color="blue"
        fullWidth
        mt="md"
        radius="md"
        onClick={() => setExpanded((e) => !e)}
      >
        Show Details
      </Button>

      <Collapse in={expanded}>
        <Paper
          radius="sm"
          style={{
            backgroundColor: "#1e2125",
            padding: "20px",
            color: "#ffffff",
            margin: "10px 0",
          }}
        >
          <Text size="xl" style={{ marginBottom: "15px" }}>
            Order Details
          </Text>
          <Divider style={{ margin: "10px 0", backgroundColor: "#ffffff30" }} />

          <Space>
            <Text
              style={{
                display: "flex",
                alignItems: "center",
                margin: "0.5rem",
              }}
            >
              <IconBrandStripe style={{ marginRight: "10px" }} />
              Stripe Payment ID :
              <strong style={{ marginLeft: "1rem" }}>
                {order.stripe_payment_id}
              </strong>
            </Text>
            <Text
              style={{
                display: "flex",
                alignItems: "center",
                margin: "0.5rem",
              }}
            >
              <IconTags style={{ marginRight: "10px" }} />
              Order Number :
              <strong style={{ marginLeft: "1rem" }}>
                {order.OrderNumber}
              </strong>
            </Text>
            <Text
              style={{
                display: "flex",
                alignItems: "center",
                margin: "0.5rem",
              }}
            >
              <IconTags style={{ marginRight: "10px" }} />
              Provider Order ID :{" "}
              <strong style={{ marginLeft: "1rem" }}>
                {order.ProviderOrderId}
              </strong>
            </Text>
            <Text
              style={{
                display: "flex",
                alignItems: "center",
                margin: "0.5rem",
              }}
            >
              <IconClockHour5 style={{ marginRight: "10px" }} />
              Order Created At :
              <strong style={{ marginLeft: "1rem" }}>
                {new Date(order.CreatedAt).toLocaleString()}
              </strong>
            </Text>
            <Text
              style={{
                display: "flex",
                alignItems: "center",
                margin: "0.5rem",
              }}
            >
              <IconClockHour5 style={{ marginRight: "10px" }} />
              Provider Created At :{" "}
              <strong style={{ marginLeft: "1rem" }}>
                {new Date(order.ProviderCreatedAt).toLocaleString()}
              </strong>
            </Text>
            <Text
              style={{
                display: "flex",
                alignItems: "center",
                margin: "0.5rem",
              }}
            >
              <IconTags style={{ marginRight: "10px" }} />
              Unique ID :{" "}
              <strong style={{ marginLeft: "1rem" }}>{order.UniqueId}</strong>
            </Text>
            <Text
              style={{
                display: "flex",
                alignItems: "center",
                margin: "0.5rem",
              }}
            >
              <IconDeviceMobile style={{ marginRight: "10px" }} />
              Customer Phone :{" "}
              <strong style={{ marginLeft: "1rem" }}>
                {order.CustomerName.DefaultAddress.Phone || "N/A"}
              </strong>
            </Text>
            <Text
              style={{
                display: "flex",
                alignItems: "center",
                margin: "0.5rem",
              }}
            >
              <IconHome2 style={{ marginRight: "10px" }} />
              Customer Address :{" "}
              <strong style={{ marginLeft: "1rem" }}>
                {order.CustomerName.DefaultAddress.Address1 +
                  ", " +
                  order.CustomerName.DefaultAddress.City +
                  ", " +
                  order.CustomerName.DefaultAddress.Province +
                  ", " +
                  order.CustomerName.DefaultAddress.Country +
                  ", " +
                  order.CustomerName.DefaultAddress.Zip || "N/A"}
              </strong>
            </Text>
          </Space>

          <Divider style={{ margin: "20px 0", backgroundColor: "#ffffff30" }} />

          <Text size="xl" style={{ marginBottom: "15px" }}>
            Products:
          </Text>
          <ul>
            {order.Products.map((product, index) => (
              <li
                key={index}
                style={{
                  marginBottom: "20px",
                  display: "flex",
                  justifyContent: "center",
                }}
              >
                <Text style={{ marginBottom: "5px", marginRight: "1rem" }}>
                  <strong>{product.Name}</strong>
                </Text>
                <Text>
                  <Badge variant="light" leftSection="💵">
                    {product.Price} {order.CustomerName.Currency}
                  </Badge>
                </Text>
              </li>
            ))}
          </ul>
        </Paper>
      </Collapse>
    </Card>
  );
};

export default OrderCard;
