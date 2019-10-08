import React from 'react'
import { Switch, Route,  } from "react-router-dom";
import Login from '../components/Login/Login'
import Portfolio from '../components/Portfolio/Portfolio'

const SiteRouter = () => {
    return (
        <Switch>
            <Route path='/' exact component={Portfolio} />
            <Route path='/login' exact component={Login} />
        </Switch>
    )
}

export default SiteRouter;
