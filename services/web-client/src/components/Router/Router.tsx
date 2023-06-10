import React from 'react';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';

const HomePage = React.lazy(() => import('pages/HomePage'));
const ProfilePage = React.lazy(() => import('pages/ProfilePage'));

const router = createBrowserRouter([
    {
        path: "/",
        element: <HomePage />
    },
    {
        path: "/profile",
        element: <ProfilePage />
    }
])

const Router: React.FunctionComponent = () => <RouterProvider router={router}/>

export default Router;