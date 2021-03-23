// import axios from "axios";

import tree from "./../tree"

const initClient = async () => { // eslint-disable-line
    try {
        tree.set("loading", true);
        tree.commit();

        // const apiUrl   = window.location.origin;
        // const urlParts = window.location.pathname.split("/");
        // const language = urlParts.length <= 2 ? "en" : urlParts[ 1 ];

        const en = "en";
        const de = "de";
        // const en = await axios({
        //     url:          `${apiUrl}/translations/${language}`,
        //     method:       "Get"
        // });

        // const de = await axios({
        //     url:          `${apiUrl}/translations/${language}`,
        //     method:       "Get"
        // });
        
        tree.set("translationPool", [{
            language:     "de",
            translations: de
        }, {
            language:     "en",
            translations: en
        }]);

        return true;
    } catch(e) {

    }
};

export default initClient;