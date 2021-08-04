import React from 'react';
import './App.css';
import Grid from '@material-ui/core/Grid';
import LocalizedStrings from 'react-localization';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';

let strings = new LocalizedStrings({
  vi: {
    selectBox_title: 'Chọn ngôn ngữ',
    homepage_title: 'Xin chào thế giới !',
    homepage_text: 'Xin chào bạn, đây là trang chủ của KFIVE'
  },
  en: {
    selectBox_title: 'Select Language',
    homepage_title: 'Hello World !',
    homepage_text: "Hi, This is KFIVE's HomePage"
  },
  ru: {
    selectBox_title: 'выберите язык ',
    homepage_title: 'Привет мир',
    homepage_text: "Здравствуйте, это домашняя страница KFIVE"
  }
});

function App() {
  const [newlanguage, setNewLanguage] = React.useState('vi')
  strings.setLanguage(newlanguage)

  const handleChange = (event) => {
    strings.setLanguage(event.target.value)
    setNewLanguage(event.target.value)
  }

  return (
    <Grid container justifyContent="center" alignItems="center">
      <Grid item xs={12} style={{ display: 'flex', justifyContent: 'center', marginTop: '40px' }}>
        <p>{strings.selectBox_title}</p>
      </Grid>
      <Grid item xs={12} style={{ display: 'flex', justifyContent: 'center' }}>
        <FormControl style={{ width: '300px' }} variant="outlined">
          <Select
            value={newlanguage}
            onChange={handleChange}
          >
            <MenuItem value={'vi'}>Tiếng việt</MenuItem>
            <MenuItem value={'en'}>English</MenuItem>
            <MenuItem value={'ru'}>русский</MenuItem>
          </Select>
        </FormControl>
      </Grid>
      <Grid item xs={12} style={{ display: 'flex', justifyContent: 'center' }}>
        <h1 className="homepage-tittle">{strings.homepage_title}</h1>
      </Grid>
      <Grid item xs={12} style={{ display: 'flex', justifyContent: 'center' }}>
        <p className="demo-paragraph">{strings.homepage_text}</p>
      </Grid>
    </Grid>
  );
}

export default App;
