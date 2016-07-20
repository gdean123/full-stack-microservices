describe("task list", function () {
    var container;

    beforeEach(function (done) {
        var link = document.createElement("link");
        link.rel = "import";
        link.href = "task_list.html";
        document.getElementsByTagName("head")[0].appendChild(link);

        container = document.createElement("div");
        container.innerHTML = "<task-list></task-list>";

        setTimeout(done, 1000);
    });

    it("displays the task list", function () {
        expect(container.innerText).toContain('Destroy the ring in Mordor');
        expect(container.innerText).toContain('Bring balance to the force');
        expect(container.innerText).toContain('Save Princess Peach... again');
        expect(container.innerText).toContain('Destroy 7 horcruxes');
        expect(container.innerText).toContain('Live long and prosper');
    })
});