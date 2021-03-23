import React from 'react';
import { Link } from 'react-router-dom';
import './NavigationBar.css';

function NavigationBar() {
  return (
    <>
      <nav className='navigationbar'>
        <div className='navigationbar-container'>
          <Link to='/' className='navigationbar-logo'>
            Hatly
          </Link>
        </div>
      </nav>
    </>
  );
}

export default NavigationBar;
