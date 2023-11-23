import { Button, Card, Container } from "react-bootstrap";
import { Navigate } from "react-router-dom";

const Home = () => {
  const token = localStorage.getItem("token");
  if (token) {
    return <Navigate to="/jobs" />;
  }
  return (
    <Container className="d-flex justify-content-center align-items-center text-center">
      <Card
        style={{ width: "25vw" }}
        className="mt-5 text-white"
        bg="dark"
        border="secondary"
      >
        <Card.Body>
          <Card.Header>
            <h2 className="fw-bold">Toped Support</h2>
          </Card.Header>
          <Card.Subtitle className="mt-2 mb-2">
            An simple app used for simulating IT support ticketing system.
          </Card.Subtitle>
          <Container className="d-flex justify-content-center align-items-center">
            <Button href="/user" variant="primary" size="lg" className="m-3">
              User
            </Button>
            <br />
            <Button href="/admin" variant="secondary" size="lg">
              Admin
            </Button>
          </Container>
        </Card.Body>
      </Card>
    </Container>
  );
};

export default Home;
