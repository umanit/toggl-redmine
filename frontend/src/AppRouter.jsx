import React from 'react';
import {HashRouter as Router, Route, Routes} from 'react-router-dom';
import Layout from './Layout';
import Home from './Home';
import Synchronise from './Synchronise';
import Configure from './Configure';

const AppRouter = () => (
  <Router>
    <Routes>
      <Route element={<Layout />} path="/">
        <Route index element={<Home />} />
        <Route path="synchroniser" element={<Synchronise />} />
        <Route path="configurer" element={<Configure />} />
      </Route>
    </Routes>
  </Router>
);

export default AppRouter
