using backend;
using backend.Settings;

const string allowedOrigins = "_myAllowSpecificOrigins";

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddCors();
builder.Services.AddSwaggerGen();
builder.Services.AddSingleton(builder.Configuration.Get<DbSettings>());
builder.Services.AddScoped<IDatabase, Database>();

builder.Services.AddCors(options => {
	options.AddPolicy(allowedOrigins, build => { // 3000 = svelte-kit dev, 3001 = svelte-kit preview
		build.WithOrigins("http://localhost:3000", "http://localhost:3001");
	});
});

var app = builder.Build();

if(app.Environment.IsDevelopment()) {
	app.UseSwagger();
	app.UseSwaggerUI();
}

app.UseHttpsRedirection();
app.UseCors(allowedOrigins);
app.UseAuthorization();
app.MapControllers();

app.Run();