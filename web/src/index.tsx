/* @refresh reload */
import { render } from 'solid-js/web';
import { Router } from '@solidjs/router';
import Layout from './layout';
import routes from './routes';
import './index.css';

const root = document.getElementById('root')!;

render(() => <Router root={Layout}>{routes}</Router>, root);
