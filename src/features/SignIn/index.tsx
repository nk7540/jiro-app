import React, {useContext, useState} from 'react';
import {StatusContext, Status} from 'services/context/StatusProvider';
import {signInWithPasswordToFirebase} from 'services/firebase';
import {SafeAreaView, View, StyleSheet} from 'react-native';
import {Input, Button, Text} from 'react-native-elements';
import {useNavigation} from '@react-navigation/native';
import {ScreenNames} from 'services/navigation';

const SignIn = () => {
  const navigation = useNavigation();
  const {setStatus} = useContext(StatusContext);
  const [input, setInput] = useState({
    email: '',
    password: '',
  });

  const signIn = async () => {
    try {
      await signInWithPasswordToFirebase(input.email, input.password);
      setStatus(Status.AUTHORIZED);
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <SafeAreaView>
      <View style={styles.form}>
        <Input placeholder="email@com" value={input.email} onChangeText={v => setInput({...input, email: v})} />
        <Input
          placeholder="password"
          value={input.password}
          onChangeText={v => setInput({...input, password: v})}
          secureTextEntry
        />
        <Button title="ログイン" titleStyle={styles.submitTitle} buttonStyle={styles.submitButton} onPress={signIn} />
        <Text style={styles.note}>
          アカウントをお持ちでないですか？
          <Text style={styles.link} onPress={() => navigation.navigate(ScreenNames.SignUp)}>
            こちら
          </Text>
          から登録できます。
        </Text>
      </View>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  form: {
    alignItems: 'center',
    marginHorizontal: 20,
  },
  note: {
    marginHorizontal: 3,
  },
  link: {
    fontWeight: 'bold',
    color: '#0047AB',
  },
  submitButton: {
    paddingHorizontal: 30,
    paddingVertical: 15,
    borderRadius: 20,
    marginBottom: 30,
  },
  submitTitle: {
    fontWeight: 'bold',
    fontSize: 20,
  },
});

export default SignIn;
