import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.OpenAPI;

@error
model NotFoundError {
  code: int32;
  message: string;
}

@error
model InternalServerError {
  code: int32;
  message: string;
}

model User {
  id: string;
  name: string;
  email: string;
}

@service({
  title: "TypeSpec practice",
})
@route("/api/v1/")
namespace API {
  @route("/users")
  namespace UserService {
    op read(@path id: string): User | NotFoundError | InternalServerError;
    op list(): User[] | InternalServerError;
  }
}
