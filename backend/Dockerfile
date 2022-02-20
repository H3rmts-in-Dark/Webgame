FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build
WORKDIR /src
COPY ["backend.csproj", "."]
RUN dotnet restore "backend.csproj"
COPY . .
RUN dotnet build "backend.csproj" -c Release -o /app/build
RUN dotnet publish "backend.csproj" -c Release -o /app/publish

FROM mcr.microsoft.com/dotnet/aspnet:6.0 AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443
COPY --from=build /app/publish .
ENTRYPOINT ["dotnet", "backend.dll"]