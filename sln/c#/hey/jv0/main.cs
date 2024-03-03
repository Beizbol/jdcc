var app = WebApplication.Create(args);
app.MapGet("/", () => "Hello World!");
app.MapGet("/hey", () => "Hey from C# and .NET 8!");
app.Run();