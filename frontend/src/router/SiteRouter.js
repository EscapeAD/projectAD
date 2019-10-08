import React from 'react'
import { Switch, Route,  } from "react-router-dom";
import { Login, Portfolio } from '../pages'

const SiteRouter = () => {
    return (
        <Switch>
            <Route path='/' exact component={Portfolio} />
            <Route path='/login' exact component={Login} />
        </Switch>
    )
}

export default SiteRouter;
