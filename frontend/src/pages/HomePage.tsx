import React, { useEffect, useState } from "react";
import { Container, Text } from "@mantine/core";
import { useNotifications } from "@mantine/notifications";
import OrderCard from "../components/OrderCard";

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
const HomePage: React.FC = () => {
  const [orders, setOrders] = useState<Order[]>([]);
  const [loading, setLoading] = useState(true);
  const notifications = useNotifications();

  useEffect(() => {
    fetch("http://localhost:8080/orders")
      .then((res) => res.json())
      .then((data) => {
        setOrders(data);
        setLoading(false);
      })
      .catch((err) => {
        setLoading(false);
        // notifications.showNotification({
        //   title: "Error fetching orders",
        //   message: err.message,
        //   color: "red",
        // });
      });
  }, [notifications]);

  if (loading) {
    return (
      <Container
        size={1000}
        style={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          height: "100vh",
        }}
      >
        {/* <Spinner /> */}
      </Container>
    );
  }

  if (!orders.length) {
    return (
      <Container size={1000}>
        <Text style={{alignItems:'center'}}>No orders found.</Text>
      </Container>
    );
  }

  return (
    <Container size={1000}>
      {orders.map((order) => (
        <OrderCard key={order.OrderNumber} order={order} />
      ))}
    </Container>
  );
};

export default HomePage;
