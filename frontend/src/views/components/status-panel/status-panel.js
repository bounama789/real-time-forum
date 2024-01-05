import { Div } from "../../elements/index.js";

export class StatusPanel{
    constructor(){
        return new Div({
            className:'status-panel',
            style:{
                flex: 1,
                display: 'flex',
                flexDirection: 'column',
                gap:"1.5rem",
                padding:"20px 0",
                backgroundColor: 'var(--bs-gray)',
                height:"100%",
            },
            children: [
            ],
        })
    }
}