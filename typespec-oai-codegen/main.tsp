import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.OpenAPI;

model GenericError<Code extends int32> {
  code: Code;
  message: string;
}

@error
model NotFoundError {
  ...GenericError<404>;
}

@error
model ConflictError {
  ...GenericError<409>;
}

@error
model InternalServerError {
  ...GenericError<500>;
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
    op create(
      @body user: OmitProperties<User, "id">,
    ): User | ConflictError | InternalServerError;

    model SendBalanceResponse {
      fromUserBalance: int32;
      toUserBalance: int32;
    }
    @error
    model InsufficientBalanceError {
      ...GenericError<4001>;
    }

    @route("/send-balance")
    op sendBalance(
      @query fromUserId: string,
      @query toUserId: string,
      amount: int32,
    ):
      | SendBalanceResponse
      | InsufficientBalanceError
      | ConflictError
      | InternalServerError;
  }
}
