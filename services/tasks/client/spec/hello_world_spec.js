describe("hello world", function() {
  beforeEach(function(done) {
    var link = document.createElement("link");
    link.rel = "import";
    link.href = "hello_world.html";
    document.getElementsByTagName("head")[0].appendChild(link);

    container = document.createElement("div");
    container.innerHTML = "<hello-world></hello-world>";
    document.body.appendChild(container);

    setTimeout(done, 1000);
  })

  it("displays hello world", function() {
    expect(document.getElementsByTagName("hello-world")[0].innerText).toEqual("Hello, world!")
  })
})
