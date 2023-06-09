def warn(*args, **kwargs):
    pass
import warnings
warnings.warn = warn
import os
os.environ['TF_CPP_MIN_LOG_LEVEL'] = '3' 
import pickle
import numpy as np
from keras.models import Sequential
from keras.layers import Bidirectional
from keras.layers import Dropout, LSTM, RepeatVector
from keras.layers import Dense, BatchNormalization
from tensorflow.keras.optimizers import Adam
from sklearn.preprocessing import StandardScaler
from sklearn.model_selection import train_test_split
import keras

Window = 5

class BILSTM:

    def __init__(self) -> None:
        self.model = self.build()

        if os.path.exists('./bilstm_model.h5'):
            self.load_model('./bilstm_model.h5')
        
        self.scr = pickle.load(open('./scr(4).pkl', 'rb'))
        self.scc = pickle.load(open('./scc(4).pkl', 'rb'))
        self.scm = pickle.load(open('./scm(4).pkl', 'rb'))


    def build(self):

        model = Sequential()
        model.add(Bidirectional(LSTM(units=5, return_sequences=True, activation="elu"), input_shape=( 5, 3)))
        model.add(Bidirectional(LSTM(units=3, activation="elu")))
        model.add(Dense(units=1, activation="elu"))
        model.build(input_shape=(None,  5, 3))

        # adamOpt = Adam(lr=0.001, beta_1=0.9, beta_2=0.999, decay=0.0, amsgrad=False)
        # model.compile(loss='mean_squared_error', optimizer=adamOpt, metrics=['mae'])

        return model
    
    def load_model(self, path):
        self.model.build(input_shape = (None, Window, 2))
        self.model.load_weights(path)
        
    def normalise_data(self, data, replicas):
        # print(data.shape)
        print(data)
        cpu = data[:, 0]
        mem = data[:, 2]
        rps = data[:, 1]/replicas


        cpu = self.scc.transform(cpu.reshape(-1, 1))

        mem = self.scm.transform(mem.reshape(-1, 1))

        rps = self.scr.transform(rps.reshape(-1, 1))

        # print(np.concatenate((cpu, rps, mem), axis=1))

        return np.concatenate((cpu, rps, mem), axis=1)

    def predict(self, X_test, replicas):
        orig = X_test[-1][0]
        X_test = self.normalise_data(X_test, replicas)
        # try:
        #     print(self.model.predict(np.array([X_test]))[0][0]/X_test[-1][0])
        # except:
        #     pass
        return self.scc.inverse_transform([self.model.predict(np.array([X_test]))[0]])
    
    def save_model(self, path):
        pickle.dump(self.model, open(path, 'wb'))

    def create_windows(self, data, windsize = Window):
        wdata = []
        ydata = []
        for i in range(windsize, len(data)-8):
            idata = []
            for j in range(windsize, 0, -1):
                idata.append(data[i - j][:2])
            wdata.append(idata)
            ydata.append(data[i+8][0])
        return np.array(wdata), np.array(ydata)

    def train(self, model, X_train):
        norm_data = self.normalise_data(X_train)

        X, y = self.create_windows(norm_data)

        X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, shuffle=False)

        self.model.fit(X_train, y_train,
                batch_size=64,
                epochs=100,
                validation_data=(X_test, y_test))
        
        self.save_model('./kmodel1.h5')

        return model
    

    
    
