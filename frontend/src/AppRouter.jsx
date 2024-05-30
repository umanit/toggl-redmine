import React from 'react';
import {HashRouter as Router, Route, Routes} from 'react-router-dom';
import Layout from './Layout/Layout';
import Home from './Home/Container';
import Synchronise from './Synchronise/Container';
import Configure from './Configure/Container';

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
