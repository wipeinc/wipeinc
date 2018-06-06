import React from 'react';
import styled from 'styled-components';
import { Link } from 'react-router-dom';


const HeroWrapper = styled.div`
  color: #282c36;
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100vh;
  position: relative;
`;

const HeroBackground = styled.div`
  display: flex;
  flex: 1;
  background-color: #f6f6f7;
  background-size: 100% auto;
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  z-index: -1;
`;

const HeroMain = styled.div`
  margin: 0 auto;
  padding-left: 4%;
  padding-right: 4%;
  width: 92%;
  @media only screen  and (min-width : 1224px) {
    width: 800px;
  }
`;

const Nav = styled.nav`
  margin: 0 auto;
  padding-left: 4%;
  padding-right: 4%;
  margin-left: auto;
  margin-right: 0;
  text-align: right;
  @media only screen  and (min-width : 1224px) {
    margin-right: 10%;
  }
`;

const NavList = styled.ul`
  list-style-type: none;
  text-transform: uppercase;
  margin: 0
  padding: 2rem 0 0 0;
  display: flex;
`;

const NavEntry = styled.li`
  flex: 1 1 auto;
  margin-left: 10px;
`;

const SignUp = styled(Link)`
  border-radius: 2px;
  color: #ef3934;
  text-align: center;
  vertical-align: middle;
  padding: 0.5rem 1rem;
  display: inline-block;
  border: 1px solid;
  cursor: pointer;
  transition: all .2s ease-in-out;
  &:focus,&:hover {
    border-color: #ef3934;
    background: #ef3934;
    color: #fff;
  }
`;

const Title = styled.h1`
  font-size: 4.4em;
`;


const Hero = () => (
  <HeroWrapper>
    <HeroBackground />
    <Nav>
      <NavList>
        <NavEntry>
          <SignUp to="/twitter/login">Sign up</SignUp>
        </NavEntry>
      </NavList>
    </Nav>
    <HeroMain>
      <Title>Wipeinc</Title>
      <h2>Tools for twitter</h2>
    </HeroMain>
  </HeroWrapper>
);

export default Hero;
