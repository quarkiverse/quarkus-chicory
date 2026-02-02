package io.quarkiverse.chicory.it;

import static io.restassured.RestAssured.given;
import static org.hamcrest.Matchers.is;

import org.junit.jupiter.api.Test;

import io.quarkus.test.junit.QuarkusTest;

@QuarkusTest
public class ChicoryResourceWithImportsTest {

    @Test
    public void testHelloEndpoint() {
        given()
                .when().get("/chicory")
                .then()
                .statusCode(200)
                .body(is("Hello chicory: " + 42));
    }
}
