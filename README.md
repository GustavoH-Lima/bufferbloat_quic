# Relatório Parte 5
### Disclaier: Professor, não consegui de jeito nenhum fazer funcionar o servidor QUIC, portanto, estou me baseando no artigo da UBER que pode ser encontrado em https://www.uber.com/en-BR/blog/employing-quic-protocol/

## 1.Qual o tempo médio de busca da página web e seu desvio padrão quando q=20 e q=100?
### No artigo não menciona

## 2.Por que você vê uma diferença nos tempos de busca de páginas da Web com buffers de roteador curtos e grandes? 
### O artigo fala que o protocolo QUIC teve uma melhora no tempo de latência em relação ao TCP principalmente para altas latências, mesmo quando o artigo usou um balancemento de cargas fornecido pela Cloud, o TCP teve uma melhora, mas o QUIC ainda foi melhor

## 4. Como o RTT relatado pelo ping varia com o tamanho da fila? Descreva a relação entre os dois.
### Em relação ao protocolo TCP, o QUIC teve uma melhora no RTT, principalmente para altos RTT, ou altas latência, mas o artigo não mencionou muito sobre a influência do Bufferbloat, apenas da conexão

## 5. Identifique e descreva duas maneiras de mitigar o problema de bufferbloat.

### O artigo menciona um melhor balanceamento de cargas para que os usuários possam se conectar no servidor mais próximo, diminuindo a latência

# Analisando dessa vez, outro artigo o "Measuring internet performances with QUIC" , temos que o TCP tende a sofrer mais de Bufferbloat em relação ao QUIC já que o TCP usa apenas um fluxo de dados. Devido ao uso de multíplas Streams do QUIC, ele apresenta mais vantagens em cenários de bufferbloat.
