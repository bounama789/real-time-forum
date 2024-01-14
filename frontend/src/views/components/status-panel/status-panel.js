import { Div } from "../../elements/index.js";
import { StatusItemView } from "./status-item-view.js";

export class StatusPanel {
    constructor() {
        return new Div({
            className: 'status-panel',
            style: {
                flex: 1,
                display: 'flex',
                flexDirection: 'column',
                gap: "1.5rem",
                padding: "20px 10px",
                height: "100%",
                backgroundColor: 'var(--bs-white)',
                boxShadow: "20px 0px 15px -23px rgba(0,0,0,0.1)",
                // top:`${getView("header").clientHeight}`
            },
            children: [
                new Div({
                    style: {
                        width: "fit-content",
                        height: "100px",
                         position: "fixed",
                         display:"flex",
                         flexDirection: "column",
                         gap: "1rem"

                    },
                    children: [
                        new StatusItemView({username:"coulou",status:"online"}),
                        new StatusItemView({username:"rehab",status:"online"}),
                        new StatusItemView({username:"jhonny",status:"offline"}),
                        new StatusItemView({username:"lampps",status:"online"}),
                        new StatusItemView({username:"kimm511",status:"offline"}),
                        new StatusItemView({username:"kuis78",status:"offline"}),


                    ],
                })
            ],
        })
    }
}