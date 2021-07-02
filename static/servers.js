'use strict';

function Servers() {
    let list = [
        ["ws://" + location.host + "/ws?name=tick", "current server"],
        ["ws://localhost:8080/ws", "local server"],
    ];

    this.getNames = () => {
        const arr = [];
        list.forEach((item) => arr.push(item[1]));
        return arr;
    };

    this.getServer = (index) => {
        if (index < 0 || index >= list.length) return undefined;
        return list[index][0];
    }

    fetch('/servers')
        .then(it => it.json())
        .then(arr => {
            const newList =
                Array.from(arr)
                    .filter(obj => !!obj.Name)
                    .map(obj => ["ws://" + location.host + "/ws?name=" + obj.Name, "Server #" + obj.Name]);

            if (!!newList) {
                list = newList;

                if (initDropdown) {
                    initDropdown();
                }
            }
        });
}

const servers = new Servers();
