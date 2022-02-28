using Microsoft.AspNetCore.Mvc.Filters;

namespace backend;

public class AddHeaderAttribute : ResultFilterAttribute {
	private readonly string _name;
	private readonly string _value;

	public AddHeaderAttribute(string name, string value) {
		_name  = name;
		_value = value;
	}

	public override void OnResultExecuting(ResultExecutingContext context) {
		context.HttpContext.Response.Headers.Add(_name, new[] {_value});
		base.OnResultExecuting(context);
	}
}

public class ServerHeader : AddHeaderAttribute {
	public ServerHeader() : base("Server", "Unix") {
	}
}