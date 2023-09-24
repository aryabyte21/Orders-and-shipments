import "./App.css";
import HomePage from "./pages/HomePage";
import { Container, Text } from "@mantine/core";

function App() {
  return (
    <div className="App">
      <header className="frosted-glass-header">
        <Container size={800}>
          <Text
            size="xl"
            fw={250}
            variant="gradient"
            gradient={{ from: "blue", to: "rgba(255, 255, 255, 1)", deg: 339 }}
          >
            Instasell Orders and Shipment Dashboard
          </Text>
        </Container>
      </header>
      <HomePage />
    </div>
  );
}

export default App;
