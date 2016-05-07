describe("hello world", function() {
  var container;

  beforeEach(function(done) {
    var link = document.createElement("link");
    link.rel = "import";
    link.href = "hello_world.html";
    document.getElementsByTagName("head")[0].appendChild(link);

    container = document.createElement("div");
    container.innerHTML = "<hello-world></hello-world>";

    setTimeout(done, 1000);
  })

  it("displays hello world", function() {
    expect(container.innerText).toEqual("Hello, world!")
  })
})
