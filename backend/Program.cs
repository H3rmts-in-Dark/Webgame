using System.Text.Json.Serialization;
using backend;
using backend.Settings;
using Microsoft.AspNetCore.Server.Kestrel.Core;

const string allowedOrigins = "_myAllowSpecificOrigins";

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers().AddJsonOptions(x => { x.JsonSerializerOptions.Converters.Add(new JsonStringEnumConverter()); });
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddCors();
builder.Services.AddSwaggerGen();
builder.Services.AddSingleton(builder.Configuration.Get<DbSettings>());
builder.Services.AddScoped<IDatabase, Database>();


// builder.WebHost.ConfigureKestrel((context, options) =>
// {
// 	options.ListenAnyIP(5252, listenOptions =>
// 	{
// 		listenOptions.Protocols = HttpProtocols.Http1AndHttp2AndHttp3;
// 	});
// });

builder.Services.AddCors(options => {
	options.AddPolicy(allowedOrigins, build => { // 3000 = svelte-kit dev, 3001 = svelte-kit preview
		                  build.AllowAnyMethod();
		                  build.AllowAnyOrigin();
		                  build.AllowAnyHeader();
	                  });
});

var app = builder.Build();

if(app.Environment.IsDevelopment()) {
	app.UseSwagger();
	app.UseSwaggerUI();
}

// app.UseHttpsRedirection();
app.UseCors(allowedOrigins);
app.UseAuthorization();
app.MapControllers();

app.Run();