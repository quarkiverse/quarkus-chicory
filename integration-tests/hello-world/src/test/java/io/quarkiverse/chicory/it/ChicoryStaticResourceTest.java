package io.quarkiverse.chicory.it;

import static io.restassured.RestAssured.given;
import static org.hamcrest.Matchers.is;

import org.junit.jupiter.api.Test;

import io.quarkus.test.junit.QuarkusTest;

@QuarkusTest
public class ChicoryStaticResourceTest {

    @Test
    public void testHelloEndpoint() {
        given()
                .when().get("/chicory/static")
                .then()
                .statusCode(200)
                .body(is("Hello chicory (static): " + 42));
    }
}
