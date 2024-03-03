Bun.serve({
    port: 4269,
    fetch(req) {
        const url = new URL(req.url);
        if (url.pathname === "/") return new Response("Home page!");
        if (url.pathname === "/hey") return new Response("Hey from Bun!");
        return new Response("404!");
    },
});
