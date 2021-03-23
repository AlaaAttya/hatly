import Baobab from "baobab";

var tree = new Baobab({
    history:         null,
    loading:         false,
    translationPool: [],
    translations:    Baobab.monkey({
        cursors: {
            pool:     ["translationPool"],
            language: ["language"]
        },
        get(data) {
            if(data.pool.length === 0) return null;

            return (data.pool.filter(elem => elem.language === data.language)[ 0 ] || {}).translations;
        }
    }),
    language: "en",
    init:     false
});

export default tree;
