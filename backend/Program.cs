using backend;
using backend.Settings;

const string myAllowSpecificOrigins = "_myAllowSpecificOrigins";

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddCors();
builder.Services.AddSwaggerGen();
builder.Services.AddSingleton(builder.Configuration.Get<DbSettings>());
builder.Services.AddScoped<IDatabase, Database>();

builder.Services.AddCors(options => {
	options.AddPolicy(myAllowSpecificOrigins, builder => { builder.WithOrigins("http://localhost:3000", "http://localhost:3001"); });
});

var app = builder.Build();

if(app.Environment.IsDevelopment()) {
	app.UseSwagger();
	app.UseSwaggerUI();
}

// app.UseHttpsRedirection(); // todo fix later
app.UseCors(myAllowSpecificOrigins);
app.UseAuthorization();
app.MapControllers();

app.Run();